import service from '@/utils/request'

// 获取pod list
export const getPodList = (data) => {
  return service({
    url: '/pod/getPodList',
    method: 'post',
    data
  })
}
