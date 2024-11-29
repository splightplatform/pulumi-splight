# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import errno
from setuptools import setup, find_packages
from setuptools.command.install import install
from subprocess import check_call


VERSION = "1.2.6"
def readme():
    try:
        with open('README.md', encoding='utf-8') as f:
            return f.read()
    except FileNotFoundError:
        return "splight Pulumi Package - Development Version"


setup(name='pulumi_splight',
      python_requires='>=3.8',
      version=VERSION,
      description="A Pulumi package for creating and managing Splight resources.",
      long_description=readme(),
      long_description_content_type='text/markdown',
      keywords='splightplatform splight infrastructure',
      url='https://www.splight-ai.com',
      project_urls={
          'Repository': 'https://github.com/splightplatform/pulumi-splight'
      },
      license='Apache-2.0',
      packages=find_packages(),
      package_data={
          'pulumi_splight': [
              'py.typed',
              'pulumi-plugin.json',
          ]
      },
      install_requires=[
          'parver>=0.2.1',
          'pulumi>=3.0.0,<4.0.0',
          'semver>=2.8.1',
          'typing-extensions>=4.11,<5; python_version < "3.11"'
      ],
      zip_safe=False)
