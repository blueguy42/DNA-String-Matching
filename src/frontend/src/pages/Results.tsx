import { FC, useState } from "react";

const axios = require('axios');

type Props = {
    className?: string;
    children: React.ReactNode; 
  };

const QueryResult: FC<Props> = ({children}) => {
    return (
        <div className="flex flex-col rounded-2xl bg-gray-700 py-3 shadow-md mx-20 my-2">
            <div className="ml-8">
                <p className="text-left">{children}</p>
            </div>
        </div>
    );
};

async function getResults(dateinput : string, diseaseinput : string) {
    try {
        const response = await axios.post('https://backend-bonek-dna.herokuapp.com/history', {date : dateinput, diseasename : diseaseinput});
        return response.data;
    } catch (error) {
        return [];
    }
}

const Results = () => {
    document.title = "Results | BONEK DNA Tester";
    
    const [queryArray, setqueryArray] = useState([]);

    async function searchQuery(query : string) {
        var hasilQuery = [];
        // format date
        if (/^\d{4}-\d{2}-\d{2}$/.test(query)) {
            hasilQuery = await getResults(query, "");

        // format date disease_name
        } else if (/^\d{4}-\d{2}-\d{2}\s/.test(query)) {
            const date = query.slice(0, 10);
            const name = query.slice(11);

            hasilQuery = await getResults(date, name);

        // format disease_name
        } else {
            hasilQuery = await getResults("", query);
        }
        if (hasilQuery.records) {
            if (hasilQuery.records.length) {
                alert ("Found " + hasilQuery.records.length + " result(s)!");
            } else {
                alert ("No results found!");
            }
        } else {
            alert ("No results found!");
        }
        
        setqueryArray(hasilQuery.records);
    }

    return (
        <>
        <div className="h=[100vh] overflow-auto">
        <div className="flex flex-col rounded-2xl bg-gray-800 shadow-md mx-64 my-10">
            <form onSubmit={(e) => {
                e.stopPropagation();
                e.preventDefault();
                searchQuery((document.getElementById("searchquery") as HTMLInputElement).value);
                (document.getElementById("searchquery") as HTMLInputElement).value = "";
                }} >
            <div className="flex flex-col lg:grid grid-cols-1 items-center my-12">
                <h1>Results</h1>
                <div className="my-12 mx-28 flex">
                    <input id="searchquery" type="text" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full mx-4 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Search query..." required/>
                    <button type="submit" className="bg-gradient-to-br w-min from-purple-600 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-14 py-2.5 text-center">Search</button>
                
                </div>
                {/* { resultBubble } */}
                {queryArray ? (queryArray.map((item,index)=>{
                    return <QueryResult key={index}>{index+1}. {item['date']} - {item['name']} - {item['disease']} - {item['similarity']}% - {item['result'] ? "True" : "False"}</QueryResult>
                })): <></>}
            </div>
            </form>
        </div>
        </div>
        </>
    );
};

export default Results;