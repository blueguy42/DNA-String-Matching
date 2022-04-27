import { useState } from "react";

const axios = require('axios');

async function addTest(patientname : string, sequencePatient : string, diseaseName : string) {
    const response = await axios.post('/predict', {name: patientname, dna: sequencePatient, diseasename : diseaseName});
    return response.data;
}

async function getDiseases() {
    try {
        const response = await axios.get('/get');
        return Array.from(response.data.names);
    } catch (error) {
        return [];
    }
}

const DNATest = () => {
    document.title = "DNA Test | BONEK DNA Tester";

    const [dnaseq, setDnaseq] = useState('');
    const [success, setSuccess] = useState(false);

    const setDNAString = (e : React.ChangeEvent<HTMLInputElement>) => {
        if (e.target.files && e.target.files[0]) {
            const reader = new FileReader();
            reader.onload = async ()=>{
                const file = reader.result;
                setDnaseq(file ? file.toString() : '');
            }
            reader.readAsText(e.target.files[0]);     
        }
    };

    async function addTestInput(patientName : string, dnaSeq : string, disease : string) {
        const diseases = await getDiseases();
        if (patientName !== "" && dnaSeq !== "" && disease !== "") {
            console.log(diseases);
            console.log(disease);
            if (diseases.includes(disease) === true) {
                if (/^[AGCT]*$/.test(dnaSeq)) {
                    addTest(patientName, dnaSeq, disease);
                    alert("Calculating results!");
                } else {
                    alert("Error! Make sure DNA sequence only contains characters AGCT!");
                }
            } else {
                alert("Error! Disease name not found!");
            }
        } else {
            alert("Error! Make sure to fill in the disease name, DNA sequence, and disease prediction!");
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
                addTestInput((document.getElementById("name") as HTMLInputElement).value, dnaseq, (document.getElementById("disease") as HTMLInputElement).value);
                (document.getElementById("name") as HTMLInputElement).value = "";
                (document.getElementById("dnasequence") as HTMLInputElement).value = "";
                (document.getElementById("disease") as HTMLInputElement).value = "";
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
                        <input id="dnasequence" type="file" className="block w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 p-2.5 cursor-pointer dark:text-gray-400 focus:outline-none focus:border-transparent dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400" required onChange={setDNAString}/>
                        </div>
                </div>
                <div className="row-start-3">
                    <h3>Disease Prediction:</h3>
                    
                        <div className="mb-6 my-6 mx-64">
                            <input id="disease" type="text" className="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="Disease Prediction..." required/>
                        </div>
                    
                </div>
            </div>
            <div className="flex flex-col lg:grid grid-cols-1 items-center">
                <div className="mt-8 mb-16">
                    <button type="submit" className="bg-gradient-to-br w-min from-purple-600 to-blue-500 hover:bg-gradient-to-bl focus:ring-4 focus:outline-none focus:ring-blue-300 dark:focus:ring-blue-800 font-medium rounded-lg text-sm px-28 py-2.5 text-center">Submit</button>
                </div>
            </div>
            </form>
        </div>
        </div>
        </>
    );
};

export default DNATest;