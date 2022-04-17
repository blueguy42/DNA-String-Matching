import { FC, useState } from "react";

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

const Results = () => {
    const [counter, setCounter] = useState(0);
    const [success, setSuccess] = useState(false);
    

    const [chats, setChats] = useState([<QueryResultInit>null</QueryResultInit>]);

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
        <div className="flex flex-col rounded-2xl bg-gray-800 shadow-md mx-64 my-24">
            <form onSubmit={(e) => {
                e.stopPropagation();
                e.preventDefault();
                setCounter(counter+1);
                changeSuccess((document.getElementById("searchquery") as HTMLInputElement).value);
                setChats([...chats, <QueryResult>{counter + ". " + (document.getElementById("searchquery") as HTMLInputElement).value}</QueryResult>]);

                (document.getElementById("searchquery") as HTMLInputElement).value = "";
                }} >
            <div className="flex flex-col lg:grid grid-cols-1 items-center my-12">
                <h1>Results</h1>
                <div className="my-6 mx-28 flex">
                    <input id="searchquery" type="text" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full mx-4 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Search query..." required/>
                    <button type="submit" className="bg-gradient-to-br w-min from-purple-600 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-14 py-2.5 text-center">Search</button>
                
                </div>
                <p className="mt-2 my-8">{ PlaceHolderText(counter, success) }</p>
                { chats }
            </div>
            </form>
        </div>
        </div>
        </>
    );
};

export default Results;