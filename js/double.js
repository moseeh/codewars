function doubleChar(str) {
    let newstr = ""
    for(let i = 0; i < str.length; i++) {
        newstr += str[i]
        newstr += str[i]
    }
    return newstr
}

console.log(doubleChar("moses"))