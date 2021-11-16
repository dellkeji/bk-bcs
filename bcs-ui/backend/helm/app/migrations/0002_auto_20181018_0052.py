# -*- coding: utf-8 -*-
"""
Tencent is pleased to support the open source community by making 蓝鲸智云PaaS平台社区版 (BlueKing PaaS Community
Edition) available.
Copyright (C) 2017-2021 THL A29 Limited, a Tencent company. All rights reserved.
Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.
"""
# Generated by Django 1.11.5 on 2018-10-17 16:52
from __future__ import unicode_literals

from django.db import migrations, models


class Migration(migrations.Migration):

    dependencies = [
        ('helm', '0002_auto_20181010_1432'),
        ('app', '0001_initial'),
    ]

    operations = [
        migrations.AddField(
            model_name='app',
            name='unique_ns',
            field=models.IntegerField(default=0),
        ),
        migrations.AlterUniqueTogether(
            name='app',
            unique_together=set([('namespace_id', 'chart', 'unique_ns')]),
        ),
        migrations.AlterIndexTogether(
            name='app',
            index_together=set([('namespace_id', 'chart', 'unique_ns')]),
        ),
    ]