import { useState } from "react";

const AddDisease = () => {
    document.title = "Add Disease | BONEK DNA Tester";

    const [counter, setCounter] = useState(0);
    const [success, setSuccess] = useState(false);

    // PLACEHOLDER BOOLEAN FUNCTION TO CHECK IF INPUT IS VALID
    function changeSuccess(nameinput : string, fileinput : string) {
        if (nameinput !== "" && fileinput !== "") {
            setSuccess(true);
        } else {
            setSuccess(false);
        }
    }

    function PlaceHolderText(counter: number, success: boolean) {
        if (counter === 0) {
            return <br/>;
        } else if (success === true) {
            return "Disease successfully added to database!";
        } else {
            return "Placeholder for warnings: " + counter;
        }
    }

    return (
        <>
        <div className="h=[100vh] overflow-hidden">
        <div className="flex flex-col rounded-2xl bg-gray-800 shadow-md mx-64 my-10">
            <form onSubmit={(e) => {
                e.stopPropagation();
                e.preventDefault();
                setCounter(counter+1);
                changeSuccess((document.getElementById("disease") as HTMLInputElement).value, (document.getElementById("dnasequence") as HTMLInputElement).value);
                (document.getElementById("disease") as HTMLInputElement).value = "";
                (document.getElementById("dnasequence") as HTMLInputElement).value = "";
                }} >
            <div className="flex flex-col my-12">
                <h1>Add Disease</h1>
            </div>
            <div className="flex flex-col lg:grid items-center">
                <div className="row-start-1">
                    <h3>Disease Name:</h3>
                    
                        <div className="mb-6 my-6 mx-64">
                        <input id="disease" type="text" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Disease Name..." required/>
                        </div>
                    
                </div>
                <div className="row-start-2">
                    <h3>DNA Sequence:</h3>
                        <div className="mb-6 my-6 mx-64">
                        <input id="dnasequence" type="file" className="block w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 p-2.5 cursor-pointer dark:text-gray-400 focus:outline-none focus:border-transparent dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400" required/>
                        </div>
                </div>
            </div>
            <div className="flex flex-col lg:grid grid-cols-1 items-center">
                <div className="mb-20">
                    <p className="my-2">{ PlaceHolderText(counter, success) }</p>
                    <button type="submit" className="bg-gradient-to-br w-min from-purple-600 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-28 py-2.5 text-center">Submit</button>
                </div>
            </div>
            </form>
        </div>
        </div>
        </>
    );
};

export default AddDisease;