// undefined
console.log(typeof undefined); // "undefined"
let person = undefined;
// console.log(person.mood); // TypeError!
let banderSnatch;
console.log(banderSnatch); // undefined

// null
let mimsy = null;
// console.log(mimsy.mood); // TypeError!
console.log(typeof null); // "object" (a lie!)

// booleans
console.log(typeof true); // "boolean"
console.log(typeof false); // "boolean"

// numbers
console.log(typeof 28); // "number"
console.log(typeof 3.14); // "number"
console.log(typeof -140); // "number"
console.log(0.1 + 0.2 === 0.3); // false
console.log(0.1 + 0.2 === 0.30000000000000004); // true
console.log(Number.MAX_SAFE_INTEGER); // 9007199254740991
console.log(Number.MAX_SAFE_INTEGER + 1); // 9007199254740992
console.log(Number.MAX_SAFE_INTEGER + 2); // 9007199254740992 (again!)
console.log(Number.MAX_SAFE_INTEGER + 3); // 9007199254740994
console.log(Number.MAX_SAFE_INTEGER + 4); // 9007199254740996
console.log(Number.MAX_SAFE_INTEGER + 5); // 9007199254740996 (again!)
let scale = 0;
let a = 1 / scale; // Infinity
let b = 0 / scale; // NaN
let c = -a; // -Infinity
let d = 1 / c; // -0
console.log(typeof NaN); // "number"

// BigInts
let aLot = 9007199254740991n; // n at the end makes it a BigInt!
console.log(aLot + 1n); // 9007199254740992n
console.log(aLot + 2n); // 9007199254740993n
console.log(aLot + 3n); // 9007199254740994n
console.log(aLot + 4n); // 9007199254740995n
console.log(aLot + 5n); // 9007199254740996n
console.log(typeof aLot); // "bigint"

// strings
console.log(typeof ""); // "string"
console.log(typeof ``); // "string"
let cat = "Cheshire";
console.log(cat.length); // 8
console.log(cat[0]); // "C"
console.log(cat[1]); // "h"

// symbols
let alohomora = Symbol();
console.log(typeof alohomora); // "symbol"
