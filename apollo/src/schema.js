import gql from "graphql-tag"

const typeDefs = gql`
  scalar Date

  type Query {
    odds: [Odd!]!
  }

  type Odd {
    id: ID!
    sport_key: String!
    sport_title: String!
    commence_time: Date!
    home_team: String!
    away_team: String!
    bookmakers: [Bookmaker!]
  }

  type Bookmaker {
    key: String!
    title: String!
    last_update: Date!
    markets: [Market!]
  }

  type Market {
    key: String!
    last_update: Date!
    outcomes: [Outcome]
  }

  type Outcome {
    name: String!
    price: Float!
  }
`

export default typeDefs
