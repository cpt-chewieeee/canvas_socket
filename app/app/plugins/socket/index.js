
module.exports = function setup (opt, imports, register) {
  register(null, {
    socket: function () {
 
      const socket = new Ws('127.0.0.1:8080/echo')
      socket.OnConnect(function () {
        console.log('socket connected connected')
        socket.Emit("chat", "hello world")
      })
      
      return socket
    }
  })
}