console.log(typeof 2); // "number"
console.log(typeof "hello"); // "string"
console.log(typeof undefined); // "undefined"

console.log(typeof {}); // "object"
console.log(typeof []); // "object"
console.log(typeof new Date()); // "object"
console.log(typeof /(hello|goodbye)/); // "object"

console.log(typeof ((x) => x * 2)); // "function"
