class Poui < Formula
  desc "Open your development web consoles with a single command"
  homepage "https://github.com/hong-yi/poui"
  url "https://github.com/hong-yi/poui.git",
    tag: "v0.1.0",
    revision: "36f67961af8bbe8992632bafc62521e4b71a4433"
  license "MIT"
  head "https://github.com/hong-yi/poui.git", branch: "main"

  livecheck do
    url :stable
    strategy :github_latest
  end

  bottle do
    sha256 cellar: :any_skip_relocation, arm64_sonoma: "87aeb068729a504f4218906b5b35365f5bc849fa1a5a0b7d30dbf7ad4adbd621"
    sha256 cellar: :any_skip_relocation, x86_64_linux: "46ea343aeb4055bfab7e6bb9f3c26d251e4aecff8aa8acbbc59eb12e5c580aed"
  end

  depends_on "go" => :build

  def install
    system "go", "build", *std_go_args, "-ldflags", "-s -w -X main.version=#{version}"
  end
end
