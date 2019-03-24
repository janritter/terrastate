# Documentation: https://docs.brew.sh/Formula-Cookbook
#                https://www.rubydoc.info/github/Homebrew/brew/master/Formula

class Terrastate < Formula
  desc "Tool to manage multiple state backends in Terraform - Allows Multi account setups"
  url "https://github.com/janritter/terrastate/releases/download/1.2.0/darwin_amd64_terrastate"
  sha256 "493e0f9e15afc67d5f4a6c699ef540d69846512f34feea2e1639926c2f1dbd0a"

  def install
    bin.install "darwin_amd64_terrastate"
    mv bin/"darwin_amd64_terrastate", bin/"terrastate"
  end

  test do
    system "false"
  end
end