# TODO

- [ ] Investigate compiling .js files with an esbuild rule in Bazel.
  - This would involve:
    - Adding a `bazel_dep` for `rules_nodejs` in `MODULE.bazel`.
    - Creating `esbuild` rules in the `BUILD.bazel` files for the `auth` and `firestore` packages.
    - This would allow us to remove the `npm run build` step from our workflow and have Bazel manage the bundling of our JavaScript files.
