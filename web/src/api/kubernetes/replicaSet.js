import service from '@/utils/request'

// 获取pod list
export const getReplicaSetList = (data) => {
  return service({
    url: '/replicaSet/getReplicaSetList',
    method: 'post',
    data
  })
}
