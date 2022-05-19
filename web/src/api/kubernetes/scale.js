import service from '@/utils/request'

// deployment replicasets 伸缩
export const scale = (data) => {
  return service({
    url: '/scale',
    method: 'post',
    data
  })
}
