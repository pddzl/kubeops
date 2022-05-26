import service from '@/utils/request'

// 创建resource
export const createResource = (data) => {
  return service({
    url: '/resource/createResource',
    method: 'post',
    data
  })
}
