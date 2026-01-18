/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

import { IPluginConfig } from '@/types';

import Icon from './assets/icon.svg?react';

export const TestrailConfig: IPluginConfig = {
  plugin: 'testrail',
  name: 'TestRail',
  icon: ({ color }) => <Icon fill={color} />,
  isBeta: true,
  sort: 18,
  connection: {
    docLink: 'https://devlake.apache.org/docs/Configuration/TestRail',
    initialValues: {
      endpoint: 'https://your-domain.testrail.io/',
    },
    fields: [
      'name',
      {
        key: 'endpoint',
        subLabel: 'Provide the TestRail instance URL (e.g., https://your-domain.testrail.io/)',
      },
      {
        key: 'username',
        label: 'Email / Username',
        placeholder: 'Enter your TestRail email or username',
      },
      {
        key: 'password',
        label: 'Password / API Key',
        type: 'password',
        placeholder: 'Enter your TestRail password or API key',
        subLabel: 'It is recommended to use an API Key instead of your login password.',
      },
      'proxy',
      {
        key: 'rateLimitPerHour',
        subLabel:
          'By default, DevLake will not limit API requests per hour. But you can set a number if you want to.',
        learnMore: 'https://devlake.apache.org/docs/Configuration/TestRail/#rate-limit-api-requests-per-hour',
        externalInfo: 'TestRail API has its own rate limits depending on your plan.',
        defaultValue: 10000,
      },
    ],
  },
  dataScope: {
    title: 'Projects',
  },
  scopeConfig: {
    entities: ['TEST', 'TESTCASE', 'TESTRESULT'],
    transformation: {},
  },
};
