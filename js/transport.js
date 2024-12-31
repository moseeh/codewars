function rentalCarCost(d) {
    if (d > 6) {
      return d*40 - 50
    } else if (d > 2) {
      return d*40 -20
    }
    return d*40
}

console.log(rentalCarCost(2))