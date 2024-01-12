let original = { metadata: { title: "My Title" } };
let copy = {
  metadata: original.metadata,
};
copy.metadata.title = "Copy of " + original.metadata.title;
console.log("original", original); // "Copy of My Title"
