import glob
import os
from tqdm import tqdm

REPLACEMENTS = [
    (
        "github.com/ethereum/go-ethereum",
        "github.com/tenderly/nitro/go-ethereum",
    ),
    (
        "github.com/tenderly/nitro/go-ethereum/internal",
        "github.com/tenderly/nitro/go-ethereum/notinternal",
    ),
    (
        "github.com/offchainlabs/nitro",
        "github.com/tenderly/nitro",
    ),
]

# REPLACEMENTS = [
#     (
#         "github.com/tenderly/nitro",
#         "github.com/offchainlabs/nitro",
#     ),
#     (
#         "github.com/tenderly/nitro/go-ethereum",
#         "github.com/ethereum/go-ethereum",
#     ),
#     (
#         "github.com/tenderly/nitro/go-ethereum/notinternal",
#         "github.com/tenderly/nitro/go-ethereum/internal",
#     ),
# ]

for filename in tqdm(glob.glob("./**", recursive=True)):
    if os.path.isdir(filename):
        continue

    if filename.endswith(".py"):
        continue

    with open(filename, "rb") as f:
        content = f.read()

    newContent = content

    for old, new in REPLACEMENTS:
        newContent = newContent.replace(old.encode(), new.encode())

    if newContent != content:
        with open(filename, "wb") as f:
            f.write(newContent)
