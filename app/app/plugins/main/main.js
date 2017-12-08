module.exports = function setup (options, imports, register) {
  const app = {}
  app.demo = imports.demo()
  app.demo.init()
  register(null, {
    main: null
  })
}