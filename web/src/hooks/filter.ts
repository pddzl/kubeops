type status = "True" | "Unknown"

const nodeStatusType: Record<status, string> = {
  True: "success",
  Unknown: "danger"
}

export const nodeStatusTypeFilter = (status: status): string => {
  return nodeStatusType[status] || "info"
}

const nodeStatus: Record<status, string> = {
  True: "Ready",
  Unknown: "NotReady"
}

export const nodeStatusFilter = (status: status) => {
  return nodeStatus[status] || "Unknown"
}
