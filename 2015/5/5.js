const fs = require('fs');

const input = fs.readFileSync('input').toString().split("\n");

const p1 = input.filter(str =>
    str.match(/[aeiou]/) && str.match(/[aeiou]/).length > 2
    && /(\w)\1/.test(str)
    && !/(ab)|(cd)|(pq)|(xy)/g.test(str)
).length
const p2 = input.filter(str => /(\w)\w\1/.test(str) && /(\w\w).*?\1/.test(str)).length

console.log(p1,p2)
