let human = {
  teeth: 32,
};

let gwen = {
  __proto__: human,
  age: 19,
};

console.log(human.age); // undefined
console.log(gwen.age); // 19

console.log(human.teeth); // 32
console.log(gwen.teeth); // 32

console.log(human.tail); // undefined
console.log(gwen.tail); // undefined

// Classes
class Spiderman {
  lookOut() {
    console.log("My Spider-Sense is tingling 1.");
  }
}

let miles1 = new Spiderman();
miles1.lookOut();

// Prototypes
let SpidermanPrototype = {
  lookOut() {
    console.log("My Spider-Sense is tingling 2.");
  },
};

let miles2 = { __proto__: SpidermanPrototype };
miles2.lookOut();
