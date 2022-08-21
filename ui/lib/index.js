import {
  LitElement,
  css,
  html,
  render,
  repeat,
  classMap,
} from "https://cdn.jsdelivr.net/gh/lit/dist@2/all/lit-all.min.js";
const baseURL = "http://localhost:8080/api/v1";
const request = (path, body, options = {}) => {
  let method = "GET";
  if (body) {
    method = "POST";
  }
  let data = {
    method,
    ...options,
  };
  if (body) {
    data.body = JSON.stringify(body);
  }
  return fetch(`${baseURL}/${path}`, data).then((r) => r.json());
};

const todoItem = ({ Title, Description, Id, Status = 0 }, { onItemClick }) => {
  const styles = { item: true, deleted: !!Status };
  return html`
    <div class=${classMap(styles)} @click=${() => onItemClick({ Id })}>
      <p>${Id}. ${Title} ${Description}</p>
    </div>
  `;
};
const todoList = (
  todos = [],
  { onItemClick, onInputChange, onAddNew, newItem } = {}
) => {
  return html`
    <div>
      <form class="box" @submit=${onAddNew}>
        <input
          class="input"
          .value=${newItem}
          placeholder="What is in your mind?"
          @change=${onInputChange}
        /><button class="button is-primary" @click=${onAddNew}>Add item</button>
      </form>
      ${repeat(
        todos,
        (todo) => todo.id,
        (todo) => todoItem(todo, { onItemClick })
      )}
    </div>
  `;
};
class AppElm extends LitElement {
  static properties = {
    todos: { state: [] },
    newItem: { state: "" },
  };
  static get styles() {
    const { cssRules } = document.styleSheets[0];
    const globalStyle = css([
      Object.values(cssRules)
        .map((rule) => rule.cssText)
        .join("\n"),
    ]);
    return [
      globalStyle,
      css`
        .container {
          width: 50%;
          margin: 0 auto;
        }
        h2 {
          text-align: center;
        }
        form {
          margin-bottom: 20px;
        }
        .item {
          border: 1px solid #d3d3d3;
          border-radius: 4px;
          margin-bottom: 10px;
          padding: 10px;
          cursor: pointer;
        }
        .item.deleted {
          background-color: #d7d7d7;
          color: red;
        }
      `,
    ];
  }

  constructor() {
    super();
    this.todos = [];
    this.newItem = "";
    this.updateComplete.then(() => {
      this.onfetchAll();
    });
  }
  onfetchAll() {
    request("todos/").then(({ data }) => {
      this.todos = data;
    });
  }

  onItemClick({ Id }) {
    const todo = this.todos.find((todo) => todo.Id == Id);
    if (todo !== -1) {
      request(`todos/${Id}`, null, { method: "POST" }).then((x) => {
        this.todos = this.todos.map((todo) => {
          if (todo.Id == Id) {
            todo.Status = 1;
          }
          return todo;
        });
      });
    }
  }
  onAddNew(e) {
    e.preventDefault();
    const item = { Title: this.newItem };
    request("todos/", item).then(({ data }) => {
      this.todos = [data, ...this.todos];
      this.newItem = "";
    });
  }
  onInputChange(e) {
    this.newItem = e.target.value;
  }
  render() {
    return html`
      <div class="container">
        <h2 class="title">Todo - DApp</h1>
        ${todoList(this.todos, {
          onItemClick: this.onItemClick.bind(this),
          onAddNew: this.onAddNew.bind(this),
          onInputChange: this.onInputChange.bind(this),
          newItem: this.newItem,
        })}
      </div>
    `;
  }
}

customElements.define("d-app", AppElm);

function App() {
  return html`<div><d-app></d-app></div>`;
}
render(App(), document.body);
