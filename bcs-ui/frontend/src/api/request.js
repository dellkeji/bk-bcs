/*
* Tencent is pleased to support the open source community by making
* 蓝鲸智云PaaS平台社区版 (BlueKing PaaS Community Edition) available.
*
* Copyright (C) 2021 THL A29 Limited, a Tencent company.  All rights reserved.
*
* 蓝鲸智云PaaS平台社区版 (BlueKing PaaS Community Edition) is licensed under the MIT License.
*
* License for 蓝鲸智云PaaS平台社区版 (BlueKing PaaS Community Edition):
*
* ---------------------------------------------------
* Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
* documentation files (the "Software"), to deal in the Software without restriction, including without limitation
* the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and
* to permit persons to whom the Software is furnished to do so, subject to the following conditions:
*
* The above copyright notice and this permission notice shall be included in all copies or substantial portions of
* the Software.
*
* THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO
* THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
* AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF
* CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS
* IN THE SOFTWARE.
*/
import http from '@/api';
import { json2Query } from '@/common/util';
import router from '@/router';
import store from '@/store';

const methodsWithoutData = ['get', 'head', 'options', 'delete'];
const defaultConfig = { needRes: false };

export const parseUrl = (url, params = {}) => {
  // 全局URL变量替换
  const { currentRoute } = router;
  const variableData = {
    $projectId: currentRoute.params.projectId,
    $clusterId: store.state.curClusterId || currentRoute.query.cluster_id || currentRoute.params.cluster_id,
  };
  Object.keys(params).forEach((key) => {
    // 自定义url变量
    if (key.indexOf('$') === 0) {
      variableData[key] = params[key];
    }
  });
  let newUrl = `${/(http|https):\/\/([\w.]+\/?)\S*/.test(url) ? url : `${DEVOPS_BCS_API_URL}${url}`}`;
  Object.keys(variableData).forEach((key) => {
    if (!variableData[key]) {
      // console.warn(`路由变量未配置${key}`)
      // 去除后面的路径符号
      newUrl = newUrl.replace(`/${key}`, '');
    } else {
      newUrl = newUrl.replace(new RegExp(`\\${key}`, 'g'), variableData[key]);
    }
    delete params[key];
  });
  return newUrl;
};

export const request = (method, url) => (params = {}, config = {}) => {
  const reqMethod = method.toLowerCase();
  const reqConfig = Object.assign({}, defaultConfig, config);

  let newUrl = parseUrl(url, params);
  let req = null;
  if (methodsWithoutData.includes(reqMethod)) {
    const query = json2Query(params, '');
    if (query) {
      newUrl += `?${query}`;
    }
    req = http[reqMethod](newUrl, null, reqConfig);
  } else {
    req = http[reqMethod](newUrl, params, reqConfig);
  }
  return req.then((res) => {
    if (reqConfig.needRes) return Promise.resolve(res);

    return Promise.resolve(res.data);
  }).catch((err) => {
    console.error('request error', err);
    return Promise.reject(err);
  });
};

export default request;
