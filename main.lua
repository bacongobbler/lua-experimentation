f = formula.new{}

version = "0.4.0"

f:name("pack-repo")
f:description("The Draft pack repository plugin")
f:homepage("https://github.com/Azure/draft-pack-repo")
f:version(version)
f:url("https://azuredraft.blob.core.windows.net/draft/draft-v" .. version .. "-darwin-amd64.tar.gz")
f:sha256("5caa5cc89d81f193615e3ad55f2c08be59052c3309f7c37d0ed0136d54b82228")

return f
