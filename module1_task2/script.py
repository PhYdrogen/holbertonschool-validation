#!/usr/bin/python3
import os

doc = "content/posts/" + os.environ.get("POST_NAME") + ".md"
name = 'title: "' + os.environ.get("POST_TITLE") + '"\n'

with open(doc, "r") as f:
    toutsauf1 = f.readlines()[2:5]
    f.close()
with open(doc, "w") as f:
    f.writelines(["---\n", name])
    f.writelines(toutsauf1)
