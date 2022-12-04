/*
Day 01...and more...
*/

function readInputFileFactory(callback) {
    return function (e) {
        var file = e.target.files[0];
        if (!file) {
            return;
        }
        var reader = new FileReader();
        reader.onload = function (e) {
            var contents = e.target.result;
            callback(contents);
        };
        reader.readAsText(file);
    }
}

function addItem(olul, txt) {
    var li = document.createElement('li');
    txt = document.createTextNode(txt);
    li.appendChild(txt);
    olul.appendChild(li);
    return li;
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

    addItem(infoElement, 'There are ' + elves.length.toString() + ' elves.');

    var maxElf = 0;
    var elfTotals = new Array();
    for (elf of elves) {
        thisElf = elf.reduce((a, b) => a + b, 0);
        elfTotals.push(thisElf);

        maxElf = Math.max(thisElf, maxElf);
    }
    outputFirst.textContent = maxElf;

    // second half...
    elfTotals.sort((a, b) => b - a);  // reverse sort
    outputSecond.textContent = (elfTotals[0] + elfTotals[1] + elfTotals[2]).toString();
}

function solveDay02(input) {
    var outputFirst = document.getElementById("day02-output-1");
    var outputSecond = document.getElementById("day02-output-2");
    var infoElement = document.getElementById("day02-info");

    addItem(infoElement, "WIP.");
}

document.getElementById('day01-input')
    .addEventListener('change', readInputFileFactory(solveDay01), false);

document.getElementById('day02-input')
    .addEventListener('change', readInputFileFactory(solveDay02), false);
