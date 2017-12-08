const Demo = function () {
  this.container = document.createElement('div')
  this.container.textContent = 'im in demo plugin'
}
Demo.prototype.init = function () {
  document.body.appendChild(this.container)
}


module.exports = function setup (options, imports, register) {
  const app = {}
  register(null, {
    demo: function() {
      return new Demo()
    }
  })
}