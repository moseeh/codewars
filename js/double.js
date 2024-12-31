function doubleChar(str) {
    let newstr = ""
    for(let i = 0; i < str.length; i++) {
        newstr += str[i].repeat(2)
    }
    return newstr
}

console.log(doubleChar("moses"))