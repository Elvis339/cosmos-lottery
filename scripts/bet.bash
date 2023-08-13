#!/bin/bash

names=()
addresses=()

# Check if the URL is active
if ! curl -s --head "http://0.0.0.0:26657" | grep "HTTP/1.1 200 OK" > /dev/null; then
    echo "Tendermint http://0.0.0.0:26657 is not active."
    exit 1
fi

# Check if ignite exists
if which ignite >/dev/null 2>&1; then
    if ! which cosmos-lotteryd >/dev/null 2>&1; then
      echo "Run ignite chain serve"
      exit 1
    else
          # Read the output line by line
          while IFS= read -r line; do
              if [[ $line =~ "name: " ]]; then
                  names+=("$(echo "$line" | awk -F': ' '{print $2}')")
              elif [[ $line =~ "- address: " ]]; then
                  addresses+=("$(echo "$line" | awk -F': ' '{print $2}')")
              fi
          done < <(cosmos-lotteryd keys list)

          for count in {0..0}; do
            echo "--- Count ${count} ---"
              for index in "${!names[@]}"; do
                  bet=${names[index]#client}
                  if [ "$bet" == "alice" ]; then
                      bet=$(( RANDOM % 100 + 1 ))
                  fi
                  addr=${addresses[index]}
#                  sleep $(( RANDOM % 10 + 1 ))
                   sleep 2

                  echo "${addr} placed a bet=${bet}"
                  cosmos-lotteryd tx lottery place-bet $bet --from $addr --yes > /dev/null
              done
          done
          echo 'Done'
          exit 0
    fi
else
    echo "ignite is not installed."
    exit 1
fi
