import "./polyfill_performance.js"
import "./wasm_exec.js"
import wasm from '../../bin/worker.wasm'
import { comment_on_issue, close_issue } from "./github.mjs"

const go = new Go()
const load = WebAssembly
.instantiate(wasm, go.importObject)
.then((instance) => {
  go.run(instance)
  return instance
})
 
const processRequest = async (event) => {
  if(event.request.method === "POST") {
    await load
    try {
      const body = await event.request.json()
      const response = await fetch(body["Link"])
      const html = await response.text()
      const results = await exec_worker(html, body["Query"])
      for(const result of results) {
        await comment_on_issue(body["Issue"], result)
      }
      await close_issue(body["Issue"])
      return new Response("Function activated!")
    } catch (error) {
      console.log(error)
      return new Response("Invalid URL supplied in request body...")
    }
  }
  return new Response(`Method Not Supported: ${event.request.method}`)
}

addEventListener("fetch", event => event.respondWith(processRequest(event)))