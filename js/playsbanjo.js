function areYouPlayingBanjo(name) {
    if (typeof name !== 'string') {
        throw new Error('Input must be a string');
    }
    if(name[0].toLowerCase() ==='r'){
     return name + " plays banjo";
    }
    return name + " does not play banjo";
}

console.log(areYouPlayingBanjo("Roses"))