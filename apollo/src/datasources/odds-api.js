import { RESTDataSource } from "@apollo/datasource-rest"

const myAPIKey = "fa5025a819281e91a5e3cadc2045d1ae"

class OddsAPI extends RESTDataSource {
  baseURL = "https://api.the-odds-api.com/"

  getOdds() {
    return this.get(
      `v4/sports/upcoming/odds/?regions=us&markets=h2h&apiKey=${myAPIKey}`
    )
  }
}

export default OddsAPI
