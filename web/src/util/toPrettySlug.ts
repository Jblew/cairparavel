import getSlug from 'speakingurl'

export function toPrettySlug(name: string) {
  return getSlug(name, { truncate: 25 })
}
