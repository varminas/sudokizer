$('input').on('keypress', function (e) {
    return e.metaKey || // cmd/ctrl
        e.which <= 0 || // arrow keys
        e.which == 8 || // delete key
        /^[1-9]{1}$/.test(String.fromCharCode(e.which)); // numbers
});

const borderStyle = 'solid 3px black';
tds = $('.sudoku-table td');
for (i = 0; i < tds.length; i++) {
    if (i % 3 === 0) {
        tds[i].style = 'border-left:' + borderStyle;
    }
      if ((i + 1) % 9 === 0) {
        tds[i].style = 'border-right:' + borderStyle;
      }
}

trs = $('.sudoku-table tr');
for (i = 0; i < trs.length; i++) {
    if ((i + 1) % 3 === 0) {
        trs[i].style = 'border-bottom:' + borderStyle;
    }
    if (i % 3 === 0) {
        trs[i].style = 'border-top:' + borderStyle;
    }
}