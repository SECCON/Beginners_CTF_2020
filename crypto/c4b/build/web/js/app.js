const app = new Vue({
  el: '#app',
  data: {
    provider: null,
    network: null,
    source: '',
    manager: null,
    challenge_address: '',
    error: '',
    flag: '',
    check: false
  },
  async mounted() {
    if (window.ethereum) {
      this.provider = new Web3(ethereum);
    } else {
      return;
    }

    this.network = await this.provider.eth.net.getNetworkType();
    if (this.network !== 'ropsten') {
      this.error = 'Switch to Ropsten Network';
    }


    fetch('/problem')
      .then(resp => resp.json())
      .then(resp => {
        this.source = resp.contract_source;
        const m_abi = JSON.parse(resp.abi.manager);
        const c_abi = JSON.parse(resp.abi.challenge);
        this.manager = new this.provider.eth.Contract(m_abi, resp.deploy_address);

        if (this.provider.givenProvider.selectedAddress) {
          this.manager.events.ChallengeDeployed({
            filter: {
              player: this.provider.givenProvider.selectedAddress
            }
          }, (error, event) => {
            if (error) {
              console.error(error);
              return;
            }
            this.error = '';
            this.challenge_address = event.returnValues.challenge;

            const challenge = new this.provider.eth.Contract(c_abi, this.challenge_address);
            challenge.events.CheckPassed({
              filter: {
                player: this.provider.givenProvider.selectedAddress
              }
            }, (error, event) => {
              this.check = true;
            });
          });
        }
      });

    fetch('/get_challenge', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({player: this.provider.givenProvider.selectedAddress})
    })
    .then(resp => resp.json())
    .then(r => {
      this.challenge_address = r.challenge;
    });

  },
  methods: {
    async deploy() {
      try {
        await ethereum.enable();
      } catch {
        this.error = 'Access denied';
        return;
      }
      this.error = '';

      this.manager.methods.deploy(this.provider.givenProvider.selectedAddress).send({
        from: this.provider.givenProvider.selectedAddress,
        gas: 500000
      }, (error, txHash) => {
        if (error) {
          this.error = error;
          return;
        }
        this.error = '';
      });
    },
    async submit() {
      if (this.challenge_address == '') {
        return;
      }

      const addr = this.provider.givenProvider.selectedAddress;
      this.provider.eth.personal.sign('Did you enjoy this challenge?\n' + addr, addr, null)
      .then(signature => {
        fetch('/submit', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json'
          },
          body: JSON.stringify({
            'signature': signature,
            'address': addr
          })
        })
          .then(resp => resp.json())
          .then(result => {
            if (result.error) {
              this.error = result.error;
            } else {
              this.error = '';
              this.flag = result.flag;
            }
          });
      });
    }
  }
});

