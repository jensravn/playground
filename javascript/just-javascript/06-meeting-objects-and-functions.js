console.log(typeof {}); // "object"
console.log(typeof []); // "object"
console.log(typeof new Date()); // "object"
console.log(typeof /\d+/); // "object"
console.log(typeof Math); // "object"

let rapper = { name: "Malicious" };
rapper.name = "Malice"; // Dot notation
rapper["name"] = "No Malice"; // Bracket notation

let countDwarves = function () {
  return 7;
};
let dwarves = countDwarves;
console.log(dwarves);
