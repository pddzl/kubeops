export const fmtTitle = (title, now) => {
  console.log('title', title)
  console.log('now', now)
  const reg = /\$\{(.+?)\}/
  const reg_g = /\$\{(.+?)\}/g
  const result = title.match(reg_g)
  if (result) {
    result.forEach((item) => {
      const key = item.match(reg)[1]
      const value = now.params[key] || now.query[key]
      title = title.replace(item, value)
    })
  }

  const reg_title = /.*详情/
  const reg_now = /.*_detail/
  if (title.match(reg_title) && now.name.match(reg_now)) {
    return `${title} / ${now.query.name}`
  } else {
    return title
  }
}
