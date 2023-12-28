import React from "react"
import { gql, useQuery } from "@apollo/client"
import { useParams } from "react-router-dom"
import { Layout, QueryResult, ModuleDetail } from "../components"

const GET_MODULE = gql`
  query getModuleAndParentTrack($moduleId: ID!, $trackId: ID!) {
    module(id: $moduleId) {
      id
      title
      content
      videoUrl
    }
    track(id: $trackId) {
      id
      title
      modules {
        id
        title
        length
      }
    }
  }
`

const Module = () => {
  const { trackId, moduleId } = useParams()
  const { loading, error, data } = useQuery(GET_MODULE, {
    variables: { trackId, moduleId },
  })

  return (
    <Layout>
      <QueryResult error={error} loading={loading} data={data}>
        <ModuleDetail track={data?.track} module={data?.module} />
      </QueryResult>
    </Layout>
  )
}

export default Module
