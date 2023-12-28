interface WordFreqMap {
  [key: string]: number;
}

export interface WordFreq {
  word: string;
  freq: number;
}

export default function termFrequency(
  text: string,
  n: number,
  stopWords: string
) {
  // Extract words
  const words = getWords(text);

  // Load stop words
  let stopWordMap = getStopWordMap(stopWords);

  // Count word frequencies
  let wordFreqs = getWordFreqs(words, stopWordMap);

  // Sort desc freq
  wordFreqs.sort((a, b) => b.freq - a.freq);

  // Prepare result of n elements
  const topN = getTopN(wordFreqs, n);

  return topN;
}

function getWords(text: string) {
  return text
    .toLowerCase()
    .replace(/[\W_']+/gm, " ")
    .split(" ");
}

function getStopWordMap(stopWords: string) {
  let stopWordMap: { [key: string]: true } = {};
  for (const word of stopWords.split(",")) {
    stopWordMap[word] = true;
  }
  return stopWordMap;
}

function getTopN(wordFreqs: WordFreq[], n: number) {
  const c = wordFreqs.length;
  n = n < c ? n : c;

  return wordFreqs.slice(0, n);
}

function getWordFreqs(words: string[], stopWordMap: { [key: string]: true }) {
  let wordFreqMap: WordFreqMap = {};
  for (const word of words) {
    if (!stopWordMap[word]) {
      wordFreqMap[word] = wordFreqMap[word] ? wordFreqMap[word] + 1 : 1;
    }
  }
  let wordFreqs: WordFreq[] = [];
  for (let [word, freq] of Object.entries(wordFreqMap)) {
    wordFreqs.push({ word, freq });
  }
  return wordFreqs;
}
