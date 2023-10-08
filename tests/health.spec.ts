describe("GET /health", () => {
  it("Should return ok", async () => {

    const res = await fetch("http://localhost:9992/health")
    const json = await res.json()
    expect(res.status).toBe(200)
    expect(json.status).toBe('ok')
    expect(Array.isArray(json.report.customizes)).toBe(true)
    expect(Array.isArray(json.report.functions)).toBe(true)
    expect(Array.isArray(json.report.proxys)).toBe(true)
    expect(isNaN(Date.parse(json.report.time))).toBe(false)
  })
})