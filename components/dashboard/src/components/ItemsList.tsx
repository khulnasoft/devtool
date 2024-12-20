/**
 * Copyright (c) 2021 Devtool GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import ContextMenu, { ContextMenuEntry } from "./ContextMenu";

export function ItemsList(props: { children?: React.ReactNode; className?: string }) {
    return <div className={`flex flex-col space-y-2 ${props.className || ""}`}>{props.children}</div>;
}

export function Item(props: { children?: React.ReactNode; className?: string; header?: boolean; solid?: boolean }) {
    let layoutClassName = "flex flex-grow flex-row justify-between";
    // set layoutClassName to "" if className contains 'grid'
    if (props.className?.includes("grid")) {
        layoutClassName = "";
    }

    // cursor-pointer hover:bg-gray-100 dark:hover:bg-gray-700
    const solidClassName = props.solid ? "bg-gray-100 dark:bg-gray-800" : "hover:bg-gray-100 dark:hover:bg-gray-800";
    const headerClassName = "text-sm text-gray-400 border-t border-b border-gray-200 dark:border-gray-800";
    const notHeaderClassName = "rounded-xl focus:bg-kumquat-light " + solidClassName;
    return (
        <div
            className={`${layoutClassName} w-full p-3 transition ease-in-out ${
                props.header ? headerClassName : notHeaderClassName
            } ${props.className || ""}`}
        >
            {props.children}
        </div>
    );
}

export function ItemField(props: { children?: React.ReactNode; className?: string }) {
    return <div className={`flex-grow mx-1 ${props.className || ""}`}>{props.children}</div>;
}

export function ItemFieldIcon(props: { children?: React.ReactNode; className?: string }) {
    return <div className={`flex self-center w-8 ${props.className || ""}`}>{props.children}</div>;
}

export function ItemFieldContextMenu(props: {
    menuEntries: ContextMenuEntry[];
    className?: string;
    position?: "start" | "center" | "end";
    changeMenuState?: (state: boolean) => void;
}) {
    const cls = "self-" + (props.position ?? "center");

    return (
        <div
            className={`flex hover:bg-gray-200 dark:hover:bg-gray-700 rounded-md cursor-pointer min-w-8 w-8 ${cls} ${
                props.className || ""
            }`}
        >
            <ContextMenu changeMenuState={props.changeMenuState} menuEntries={props.menuEntries} />
        </div>
    );
}
