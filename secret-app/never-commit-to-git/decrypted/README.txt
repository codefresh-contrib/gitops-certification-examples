The whole point of this example is to use encrypted
secrets in Git.

Thus you should NEVER commit encrypted stuff in git.
But just for demo purposes this directory contains
the raw secrets for your convenience. This way you 
can compare what you see deployed in Kubernetes
with the contents of this directory.

In a real production application only the 
encrypted stuff MUST be committed in Git (see
the manifests folder).