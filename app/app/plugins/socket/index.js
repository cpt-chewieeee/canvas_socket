const io = require('socket.io-client')
console.log('what is ', io)

module.exports = function setup (opt, imports, register) {
  register(null, {
    socket: function () {

      const socket = io('http://localhost:3131', { transports: ['websocket'] })

      socket.on('connect', function () {
        console.log('socket connect')
        socket.emit('send', { name : 'test_name', message: 'hello_world' }, function (result) {
          console.log('sended successfully')
        })
      })
      console.log('--------->socket<---------')
      return socket
    }
  })
}