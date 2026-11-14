import { request } from '../request';

export function fetchDevicesList(params: any) {
  return request<Api.Devices.DevicesList>({ url: '/devices/list', params });
}

export function updateDevice(data: { id: number; alias: string; company?: string; department: string }) {
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

export function fetchCompanyList() {
  return request<string[]>({ url: '/companies/list' });
}

export function fetchCompanyTree() {
  return request<Api.Devices.CompanyTreeItem[]>({ url: '/companies/tree' });
}

export function addCompany(data: { name: string; note?: string }) {
  return request({
    url: '/companies/add',
    method: 'post',
    data
  });
}

export function deleteCompany(name: string) {
  return request({
    url: '/companies/delete',
    method: 'post',
    data: { name }
  });
}

export function addDepartment(data: { company_name: string; name: string }) {
  return request({
    url: '/departments/add',
    method: 'post',
    data
  });
}

export function deleteDepartment(data: { company_name: string; name: string }) {
  return request({
    url: '/departments/delete',
    method: 'post',
    data
  });
}
