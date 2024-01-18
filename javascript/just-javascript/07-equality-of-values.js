function strictEquals(a, b) {
  if (Number.isNaN(a) && Number.isNaN(b)) {
    return false;
  }
  if (
    (Object.is(a, 0) && Object.is(b, -0)) ||
    (Object.is(a, -0) && Object.is(b, 0))
  ) {
    return true;
  }
  return Object.is(a, b);
}

// special cases
console.log("NaN === NaN", NaN === NaN, strictEquals(NaN, NaN)); // NaN === NaN false false
console.log("-0 === 0", -0 === 0, strictEquals(-0, 0)); // -0 === 0 true true
console.log("0 === -0", -0 === 0, strictEquals(-0, 0)); // 0 === -0 true true

// NOT special cases
console.log("1 === 1", 1 === 1, strictEquals(1, 1)); // 1 === 1 true true
console.log("{} === {}", {} === {}, strictEquals({}, {})); // {} === {} false false
console.log("-0 === -0", -0 === -0, strictEquals(-0, -0)); // -0 === -0 true true
console.log("0 === 0", 0 === 0, strictEquals(0, 0)); // 0 === 0 true true
