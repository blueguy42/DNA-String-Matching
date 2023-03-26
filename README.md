# DNA String Matching

This program was created to fulfill **IF2211 Algorithm Strategies: Major Assignment 3** in Semester II 2021/2022. 

## Table of Contents
* [General Info](#general-information)
* [Algorithms Used](#algorithms-used)
* [Technologies Used](#technologies-used)
* [Features](#features)
* [Screenshots](#screenshots)
* [Setup](#setup)
* [Usage](#usage)
* [Project Status](#project-status)
* [Contact](#contact)


## General Information
DNA is the genetic material that determines the traits and characteristics of an individual. When someone has a genetic disorder or abnormality in their DNA, whether inherited or due to environmental factors, they may develop a particular disease. Genetic testing can anticipate such genetic diseases by examining the genetic structure within an individual's body and detecting any abnormalities. There are various types of DNA tests in the field of bioinformatics that can be performed, and one of them is DNA sequence analysis. DNA sequence analysis is a way to predict various diseases stored in a database based on the sequence of DNA, represented by the sequence of four symbols [A, G, C, T] representing the nitrogenous bases Adenine, Guanine, Cytosine, and Thymine.

DNA sequence analysis can be done using pattern matching techniques that can analyze DNA sequences quickly. This program is a DNA sequence matcher program that performs pattern matching and determines if a human DNA contains a specific disease. The program first validates the sequence of human DNA or disease from user input using Regular Expression. Then, the program applies the KMP and BM algorithms to perform string matching and determine if the disease DNA is contained within the human DNA. The program can also determine the similarity of both DNA using LCS and Gestalt Pattern Matching algorithms.


## Algorithms Used
- Knuth-Morris-Pratt (KMP) : A string matching algorithm performs pattern search in a string sequentially from left to right, similar to the Brute Force algorithm, but it has a smarter search pattern that improves the efficiency of the search by reducing the number of letter comparisons examined. This algorithm utilizes the matching of prefixes and suffixes of a string. When searching, if a mismatch is found between a letter in string S at index i and a letter in pattern P at index j, then index i in S can be shifted by the largest prefix of P[0..j-1] that is also a suffix of P[1..j-1].

- Boyer-Moore (BM) : The string matching algorithm is performed by matching the pattern from right to left. The examination of a string S starts from the beginning, but the examination of a pattern P starts from its last index, which is the length of the pattern minus one. This algorithm utilizes two techniques, namely the looking-glass technique and the character-jump technique. The looking-glass technique involves examining a pattern P in a string S by moving backward towards P, starting from the end. If the examined letter is correct, the examination continues to the left until all letters in the pattern have been examined. If the examined letter is incorrect, the character-jump technique is utilized. The character-jump technique involves shifting the index i in S if a mismatched letter is found, and returning it to the last index in P.

- Longest Common Subsequence (LCS) : Longest Common Subsequence is a problem of finding the longest subsequence of two strings and is used to determine the similarity of the two strings. Subsequence differs from substring in that it does not have to be contiguous, while substring must be contiguous, but subsequence must still be relatively ordered by its occurrence. Similarity will be determined using the Gestalt Pattern Matching formula, which utilizes the length of the LCS itself.

## Technologies Used
### Programming Language
- Golang
- SQL
- HTML
- CSS
- Typescript
### Frameworks
- React
- Tailwind

## Features
- The application can receive new disease inputs in the form of disease names and their DNA sequences (which are then added to the database).
- The application can predict whether someone has a certain disease based on their DNA sequence using the KMP and BM algorithms.
- The application has a page that displays the order of prediction results with a search column, where the search column can work as a filter. The filter can work with three cases, namely disease names only, dates only, or both.
- The application can validate DNA input using Regex.
- The application can calculate the level of similarity between a user's DNA and the disease DNA in DNA testing using the LCS algorithm.


## Screenshots
![predict.png](./img/Predict.png)
![History.png](./img/History.png)
![addDisease.png](./img/addDisease.png)

## Setup
1. Make sure to install yarn by installing node first through the link https://nodejs.org/en/download/ and running the following command:
```
npm install --global yarn
```
2. Make sure that you also have a browser, preferably Google Chrome.


## Usage
There are two ways to use the program that has been provided. The first way is by running the web application locally by executing the following command:
```
// go to src/frontend folder of this repository
cd src/frontend
// make sure that yarn has been installed
yarn start
```

Users do not need to run the backend server and initialize the database locally because the frontend, which is executed with the above command, is already connected to the backend and database in the cloud that has been deployed using Heroku.

The second and easier way is to directly access the deployed web application through the address https://bonek-dna.netlify.app/

## Project Status
Project is: **Complete**

## Contact
This project was created by group 5 **(BONEK Returns)** consisting of:
>- <a href="https://www.linkedin.com/in/ahmad-alfani-handoyo/"> Ahmad Alfani Handoyo (13520023)</a>
>- <a href="https://www.linkedin.com/in/saulsayers/?originalSubdomain=id">Saul Sayers (13520094)</a>
>- <a href="https://www.linkedin.com/in/rizky-ramadhana-putra-kusnaryanto-6037a51aa/">Rizky Ramadhana Putra Kusnaryanto (13520151)</a>
