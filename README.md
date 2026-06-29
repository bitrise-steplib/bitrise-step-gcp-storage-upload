# Google Cloud Storage Upload

[![Step changelog](https://shields.io/github/v/release/bitrise-steplib/bitrise-step-gcp-storage-upload?include_prereleases&label=changelog&color=blueviolet)](https://github.com/bitrise-steplib/bitrise-step-gcp-storage-upload/releases)

The Step uploads files to a Google Cloud Storage bucket.

❗ There's no retry logic implemented in this step at this point. A failed object upload halts the whole upload process.

<details>
<summary>Description</summary>

The Step uploads files to a Google Cloud Storage bucket.

It can handle both files and directories, uploading the contents to the specified GCS bucket with the desired access control settings.

Using the Authenticate with GCP Step beforehand is recommended to securely provide short lived GCP credentials.
</details>

## 🧩 Get started

Add this step directly to your workflow in the [Bitrise Workflow Editor](https://docs.bitrise.io/en/bitrise-ci/workflows-and-pipelines/steps/adding-steps-to-a-workflow.html).

You can also run this step directly with [Bitrise CLI](https://github.com/bitrise-io/bitrise).

## ⚙️ Configuration

<details>
<summary>Inputs</summary>

| Key | Description | Flags | Default |
| --- | --- | --- | --- |
| `path` | Path to a file or folder to be uploaded.  You can use absolute or relative paths. | required |  |
| `bucket_name` | Name of the GCS bucket to upload the file to. | required |  |
| `bucket_prefix` | Path in the GCS bucket where the file will be uploaded.  If not provided, the file will be uploaded to the root of the bucket with its original filename. |  |  |
| `access_token` | The GCP Access Token.  You can provide it directly or use the Authenticate with GCP Step. | required, sensitive | `$GOOGLE_AUTH_TOKEN` |
| `verbose` | Enable logging additional information for debugging. | required | `false` |
</details>

<details>
<summary>Outputs</summary>
There are no outputs defined in this step
</details>

## 🙋 Contributing

We welcome [pull requests](https://github.com/bitrise-steplib/bitrise-step-gcp-storage-upload/pulls) and [issues](https://github.com/bitrise-steplib/bitrise-step-gcp-storage-upload/issues) against this repository.

For pull requests, work on your changes in a forked repository and use the Bitrise CLI to [run step tests locally](https://docs.bitrise.io/en/bitrise-ci/bitrise-cli/running-your-first-local-build-with-the-cli.html).

Learn more about developing steps:

- [Create your own step](https://docs.bitrise.io/en/bitrise-ci/workflows-and-pipelines/developing-your-own-bitrise-step/developing-a-new-step.html)
