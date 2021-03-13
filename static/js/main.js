$('input').on('keypress', function(e){
    return e.metaKey || // cmd/ctrl
      e.which <= 0 || // arrow keys
      e.which == 8 || // delete key
      /^[1-9]{1}$/.test(String.fromCharCode(e.which)); // numbers
  })