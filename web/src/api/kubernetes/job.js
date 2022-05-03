import service from '@/utils/request'

// 获取job list
export const getJobList = (data) => {
  return service({
    url: '/job/getJobList',
    method: 'post',
    data
  })
}

// 获取job编排
export const getJobRaw = (data) => {
  return service({
    url: '/job/getJobRaw',
    method: 'post',
    data
  })
}

// 获取job in detail
export const getJobDetail = (data) => {
  return service({
    url: '/job/getJobDetail',
    method: 'post',
    data
  })
}

// 获取job关联pods
export const getJobPods = (data) => {
  return service({
    url: '/job/getJobPods',
    method: 'post',
    data
  })
}
