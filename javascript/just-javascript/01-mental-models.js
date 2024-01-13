// reassign primitive value
let a = 5; // a is assigned the value 5
let b = a; // b is also assigned the value 5
a = 0; // a is reassigned the value 0
console.log("a", a);
console.log("b", b);

// reassign object
let original = { metadata: { title: "My Title" } };
let copy = {
  metadata: original.metadata,
};
copy.metadata.title = "Copy of " + original.metadata.title;
console.log("original", original); // "Copy of My Title"
