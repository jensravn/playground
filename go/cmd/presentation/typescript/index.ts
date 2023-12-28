import * as fs from "fs";
import termFrequency from "./src/termFrequency";

// Read input files
const text = fs.readFileSync("../_resources/pride-and-prejudice.txt", "utf8");
const n = 25;
const stopWords = fs.readFileSync("../_resources/stop_words.txt", "utf8");

const start = new Date().getTime();
console.log("START");

// Term frequency
const mostUsedWords = termFrequency(text, n, stopWords);

const end = new Date().getTime();
console.log(`END - ${end - start}ms\n`);

// Print result
for (let i = 0; i < n; i++) {
  console.log(`${mostUsedWords[i].word} - ${mostUsedWords[i].freq}`);
}
