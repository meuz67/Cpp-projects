const fs = require('fs');
function characterCount(str) {
    return str.length;
}
function analyzeFile(fileName){
    let fileContext = fs.readFileSync(fileName, 'utf8');
    return characterCount(fileContext);
}
console.log(analyzeFile("file.txt"));