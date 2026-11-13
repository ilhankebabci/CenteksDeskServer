import { request } from '../request';

export function fetchDevicesList(params: any) {
  return request<Api.Devices.DevicesList>({ url: '/devices/list', params });
}

export function updateDevice(data: { id: number; alias: string; department: string }) {
  return request({
    url: '/devices/update',
    method: 'post',
    data
  });
}

export function deleteDevice(id: number) {
  return request({
    url: '/devices/delete',
    method: 'post',
    data: { id }
  });
}
