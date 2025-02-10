import { Gitlab } from '@gitbeaker/rest'
import {getEnvVariable} from './utils'

let gitlabApi: InstanceType<typeof Gitlab> | null = null

async function initApi(): Promise<InstanceType<typeof Gitlab>> {
  const token = getEnvVariable('GITLAB_TOKEN')
  let host = process.env['CI_SERVER_HOST'] || "https://gitlab.com"
  if (!host.startsWith("https://")) {
    host = `https://${host}`
  }
  const gitlab = new Gitlab({
    token: token,
    host: host
  })
  gitlabApi = gitlab
  return gitlab
}

export async function getGitlabApi(): Promise<InstanceType<typeof Gitlab>> {
  try {
    const api = gitlabApi
    if (!api) {
      return await initApi()
    } else {
      return api
    }
  } catch (e) {
    console.error("Could not access Gitlab API")
    throw e
  }
}