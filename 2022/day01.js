function readInputFile(e) {
    var file = e.target.files[0];
    if (!file) {
        return;
    }
    var reader = new FileReader();
    reader.onload = function(e) {
        var contents = e.target.result;
        solveDay01(contents);
    };
    reader.readAsText(file);
}


function solveDay01(input) {
    var outputFirst = document.getElementById("day01-output-1");
    var outputSecond = document.getElementById("day01-output-2");
    var infoElement = document.getElementById("day01-info");

    var elves = new Array();
    var elf = new Array();
    elves.push(elf);
    for (line of input.split('\n')) {
        if (line == "") {
            elf = new Array();
            elves.push(elf);
        } else {
            elf.push(Number.parseInt(line));
        }
    }
    // console.log(elves);

    var li = document.createElement('li');
    txt = document.createTextNode('There are ' + elves.length.toString() + ' elves.');
    li.appendChild(txt);
    infoElement.appendChild(li);

    var maxElf = 0;
    var elfTotals = new Array();
    for (elf of elves) {
        thisElf = elf.reduce((a,b)=>a+b, 0);
        elfTotals.push(thisElf);

        maxElf = Math.max(thisElf, maxElf);
    }
    outputFirst.textContent = maxElf;

    // second half...
    elfTotals.sort((a, b) => b - a);
    outputSecond.textContent = (elfTotals[0] + elfTotals[1] + elfTotals[2]).toString();
}

document.getElementById('day01-input').addEventListener('change', readInputFile, false);
