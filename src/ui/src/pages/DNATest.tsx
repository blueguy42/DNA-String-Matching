import { useState } from "react";

const axios = require('axios');

export async function getDiseases() {
    console.log("TEST");
    try {
        console.log("1---");
        const response = await axios.get('/get');
        console.log("2---");
        console.log('response: ', response);
        console.log("3---");
        return response.data;
    } catch (error) {
        console.log("4---");
        console.log('error: ', error);
        console.log("5---");
        return [];
    }
}

const DNATest = () => {
    document.title = "DNA Test | BONEK DNA Tester";

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
            return "DNA test completed!";
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
                changeSuccess((document.getElementById("name") as HTMLInputElement).value, (document.getElementById("dnasequence") as HTMLInputElement).value);
                (document.getElementById("name") as HTMLInputElement).value = "";
                (document.getElementById("dnasequence") as HTMLInputElement).value = "";
                (document.getElementById("disease") as HTMLInputElement).value = "1";
                }} >
            <div className="flex flex-col my-12">
                <h1>DNA Test</h1>
            </div>
            <div className="flex flex-col lg:grid items-center">
                <div className="row-start-1">
                    <h3>Patient Name:</h3>
                    
                        <div className="mb-6 my-6 mx-64">
                            <input id="name" type="text" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Patient Name..." required/>
                        </div>
                    
                </div>
                <div className="row-start-2">
                    <h3>DNA Sequence:</h3>
                        <div className="mb-6 my-6 mx-64">
                        <input id="dnasequence" type="file" className="block w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 p-2.5 cursor-pointer dark:text-gray-400 focus:outline-none focus:border-transparent dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400" required/>
                        </div>
                </div>
                <div className="row-start-3">
                    <h3>Disease Prediction:</h3>
                    
                        <div className="mb-6 my-6 mx-64">
                            <select id="disease" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" required>
                                <option value="1">Nanti request</option>
                                <option value="2">semua disease</option>
                                <option value="3">yang ada di database</option>
                                <option value="4">dong. biar bisa</option>
                                <option value="5">di-call di sini</option>
                            </select>

                        </div>
                    
                </div>
            </div>
            <div className="flex flex-col lg:grid grid-cols-1 items-center">
                <div>
                    <p className="my-2">{ PlaceHolderText(counter, success) }</p>
                    <button type="submit" className="bg-gradient-to-br w-min from-purple-600 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-28 py-2.5 text-center">Submit</button>
                </div>
            </div>
            </form>
            <div className="mt-6 mb-14">
                { success ? 
                <div>
                <p>Test Result:</p>
                <p>1 April 2022 - Mhs IF - HIV - 75% - False (placeholder)</p>
                </div>
                : null }
            </div> 
        </div>
        </div>
        </>
    );
};

export default DNATest;