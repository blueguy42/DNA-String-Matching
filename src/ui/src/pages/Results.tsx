import { FC, useState } from "react";

const axios = require('axios');

type Props = {
    className?: string;
    children: React.ReactNode; 
  };
  
const QueryResultInit: FC<Props> = ({children}) => {
return (<></>);
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
        console.log("date: " + dateinput + ", disease: " + diseaseinput);
        const response = await axios.get('/history', {date : dateinput, diseasename : diseaseinput});
        return response.data;
    } catch (error) {
        return [];
    }
}

const Results = () => {
    document.title = "Results | BONEK DNA Tester";

    const [counter, setCounter] = useState(0);
    const [success, setSuccess] = useState(false);

    const [query, setquery] = useState('');
    

    const [resultBubble, setresultBubble] = useState([<QueryResultInit>null</QueryResultInit>]);

    async function searchQuery(query : string) {
        // format date
        if (/^\d{4}\-\d{2}\-\d{2}$/.test(query)) {
            console.log("format date");
            
            const hasilQuery = await getResults(query, "");
            console.log(hasilQuery);

        // format date disease_name
        } else if (/^\d{4}\-\d{2}\-\d{2}\s/.test(query)) {
            console.log("format date disease_name");
            const date = query.slice(0, 10);
            const name = query.slice(11);

            const hasilQuery = await getResults(date, name);
            console.log(hasilQuery);

        // format disease_name
        } else {

            console.log("format disease_name");

            const hasilQuery = await getResults("", query);
            console.log(hasilQuery);
        }
    }

    // PLACEHOLDER BOOLEAN FUNCTION TO CHECK IF INPUT IS VALID
    function changeSuccess(nameinput : string) {
        if (nameinput !== "") {
            setSuccess(true);
        } else {
            setSuccess(false);
        }
    }

    function PlaceHolderText(counter: number, success: boolean) {
        if (counter === 0) {
            return <br/>;
        } else if (success === true) {
            return "Search successful with " + counter + " results returned!";
        } else {
            return "Placeholder for warnings: " + counter;
        }
    }

    return (
        <>
        <div className="h=[100vh] overflow-auto">
        <div className="flex flex-col rounded-2xl bg-gray-800 shadow-md mx-64 my-10">
            <form onSubmit={(e) => {
                e.stopPropagation();
                e.preventDefault();
                setCounter(counter+1);
                changeSuccess((document.getElementById("searchquery") as HTMLInputElement).value);
                searchQuery((document.getElementById("searchquery") as HTMLInputElement).value);
                setresultBubble([...resultBubble, <QueryResult>{counter+1 + ". " + (document.getElementById("searchquery") as HTMLInputElement).value}</QueryResult>]);

                (document.getElementById("searchquery") as HTMLInputElement).value = "";
                }} >
            <div className="flex flex-col lg:grid grid-cols-1 items-center my-12">
                <h1>Results</h1>
                <div className="my-12 mx-28 flex">
                    <input id="searchquery" type="text" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full mx-4 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Search query..." required/>
                    <button type="submit" className="bg-gradient-to-br w-min from-purple-600 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-14 py-2.5 text-center">Search</button>
                
                </div>
                { resultBubble }
            </div>
            </form>
        </div>
        </div>
        </>
    );
};

export default Results;